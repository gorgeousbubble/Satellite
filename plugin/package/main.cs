using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;

namespace package
{
    public partial class FormMain : Form
    {
        struct T_PackInfo
        {
            public int Number;
            public string Name;
            public string Path;
        }

        private List<T_PackInfo> m_vecPackInfo = new List<T_PackInfo>();

        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // TabPack Initial
            listView_pack.View = View.Details;
            listView_pack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_pack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_pack.Columns.Add("detials", 400, HorizontalAlignment.Center);

            comboBox_pack.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox_pack.Items.Clear();
            comboBox_pack.Items.Add("aes");
            comboBox_pack.Items.Add("3des");
            comboBox_pack.Items.Add("des");
            comboBox_pack.Items.Add("rsa");
            comboBox_pack.Items.Add("base64");
            comboBox_pack.SelectedIndex = 0;

        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }

        private void Button_pack_add_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = true;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                foreach (var i in dialog.FileNames)
                {
                    bool flag = false;
                    foreach (var j in m_vecPackInfo)
                    {
                        if(j.Path == i)
                        {
                            flag = true;
                            break;
                        }
                    }
                    if(!flag)
                    {
                        T_PackInfo packInfo = new T_PackInfo();
                        FileInfo file = new FileInfo(i);
                        packInfo.Number = m_vecPackInfo.Count + 1;
                        packInfo.Name = file.Name;
                        packInfo.Path = i;
                        m_vecPackInfo.Add(packInfo);
                    }
                }

                listView_pack.Items.Clear();
                listView_pack.BeginUpdate();
                for(int i = 0; i < m_vecPackInfo.Count; ++i)
                {
                    ListViewItem item = new ListViewItem();
                    item.Text = m_vecPackInfo[i].Number.ToString();
                    item.SubItems.Add(m_vecPackInfo[i].Name);
                    item.SubItems.Add(m_vecPackInfo[i].Path);
                    listView_pack.Items.Add(item);
                }
                listView_pack.EndUpdate();
            }

        }

        private void Button_pack_delete_Click(object sender, EventArgs e)
        {

        }
    }
}
